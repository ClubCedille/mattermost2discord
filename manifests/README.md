# Kustomize configuration

If you plan to deploy `mattermost2discord` using [Kustomize](https://kustomize.io), do not forget to also declare the following secret:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: mm2disc-secrets
  namespace: mattermost2discord
type: Opaque
data:
  DISCORD_TOKEN: '<REPLACE_ME>'
  DISCORD_CHANNEL: '<REPLACE_ME>'
  MATTERMOST_TOKEN: '<REPLACE_ME>'
  TRIGGER_WORD_MATTERMOST: '<REPLACE_ME>'
```
