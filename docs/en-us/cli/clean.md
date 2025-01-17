Ktctl Clean
---

Delete unavailing resources created by kt from kubernetes cluster. Basic usage:

```bash
ktctl clean
```

Available options:

```
--dryRun                  Only print name of resources to be deleted
--thresholdInMinus value  Length of allowed disconnection time before a unavailing shadow pod be deleted (default: 15)
--sweepLocalRoute         Also clean up local route table record created by kt
```

Key options explanation:

- The value of the `--thresholdInMinus` parameter should not be less than the default heartbeat interval of KT resources (5 minutes), otherwise normal resources in use may be deleted unexpectedly.
