package dns

import (
	"fmt"
	"github.com/alibaba/kt-connect/pkg/common"
	opt "github.com/alibaba/kt-connect/pkg/kt/command/options"
	"github.com/alibaba/kt-connect/pkg/kt/service/cluster"
	"github.com/alibaba/kt-connect/pkg/kt/util"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

const (
	resolverDir = "/etc/resolver"
	ktResolverPrefix = "kt."
	resolverComment  = "# Generated by KtConnect"
)

// SetNameServer set dns server records
func (s *Cli) SetNameServer(dnsServer string) error {
	dnsSignal := make(chan error)
	if err := util.CreateDirIfNotExist(resolverDir); err != nil {
		log.Error().Err(err).Msgf("Failed to create resolver dir")
		return err
	}
	go func() {
		namespaces, err := cluster.Ins().GetAllNamespaces()
		if err != nil {
			dnsSignal <-err
			return
		}

		preferredDnsInfo := strings.Split(dnsServer, ":")
		dnsIp := preferredDnsInfo[0]
		dnsPort := strconv.Itoa(common.StandardDnsPort)
		if len(preferredDnsInfo) > 1 {
			dnsPort = preferredDnsInfo[1]
		}

		createResolverFile("local", opt.Get().ConnectOptions.ClusterDomain, dnsIp, dnsPort)
		for _, ns := range namespaces.Items {
			createResolverFile(fmt.Sprintf("%s.local", ns.Name), ns.Name, dnsIp, dnsPort)
		}
		dnsSignal <- nil

		defer s.RestoreNameServer()
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
	}()
	return <-dnsSignal
}

func createResolverFile(postfix, domain, dnsIp, dnsPort string) {
	resolverFile := fmt.Sprintf("%s/%s%s", resolverDir, ktResolverPrefix, postfix)
	if _, err := os.Stat(resolverFile); err == nil {
		_ = os.Remove(resolverFile)
	}
	resolverContent := fmt.Sprintf("%s\ndomain %s\nnameserver %s\nport %s\n",
		resolverComment, domain, dnsIp, dnsPort)
	if err := ioutil.WriteFile(resolverFile, []byte(resolverContent), 0644); err != nil {
		log.Warn().Err(err).Msgf("Failed to create resolver file of %s", domain)
	}
}

// RestoreNameServer remove the nameservers added by ktctl
func (s *Cli) RestoreNameServer() {
	rd, _ := ioutil.ReadDir(resolverDir)
	for _, f := range rd {
		if !f.IsDir() && strings.HasPrefix(f.Name(), ktResolverPrefix) {
			if err := os.Remove(fmt.Sprintf("%s/%s", resolverDir, f.Name())); err != nil {
				log.Warn().Err(err).Msgf("Failed to remove resolver file %s", f.Name())
			}
		}
	}
}
