package pkg

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestYAMLReader(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "empty",
			input: "",
			want:  []string{},
		},
		{
			name: "list with newlines in configmap",
			input: `apiVersion: v1
items:
- apiVersion: v1
  data:
    status: |+
      Cluster-autoscaler status at 2021-11-10 00:01:27.986516659 +0000 UTC:
      Cluster-wide:
        Health:      Healthy (ready=6 unready=0 notStarted=0 longNotStarted=0 registered=6 longUnregistered=0)
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-10-26 05:20:47.050658932 +0000 UTC m=+238.651302118
        ScaleUp:     NoActivity (ready=6 registered=6)
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-11-09 08:49:17.188499976 +0000 UTC m=+1222348.789143150
        ScaleDown:   NoCandidates (candidates=0)
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-11-09 08:53:30.879288373 +0000 UTC m=+1222602.479931549

      NodeGroups:
        Name:        obscured
        Health:      Healthy (ready=1 unready=0 notStarted=0 longNotStarted=0 registered=1 longUnregistered=0 cloudProviderTarget=1 (minSize=1, maxSize=100))
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-10-26 05:20:47.050658932 +0000 UTC m=+238.651302118
        ScaleUp:     NoActivity (ready=1 cloudProviderTarget=1)
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-10-26 05:20:47.050658932 +0000 UTC m=+238.651302118
        ScaleDown:   NoCandidates (candidates=0)
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-10-26 05:20:47.050658932 +0000 UTC m=+238.651302118

        Name:        obscured
        Health:      Healthy (ready=1 unready=0 notStarted=0 longNotStarted=0 registered=1 longUnregistered=0 cloudProviderTarget=1 (minSize=1, maxSize=100))
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-10-26 05:20:47.050658932 +0000 UTC m=+238.651302118
        ScaleUp:     NoActivity (ready=1 cloudProviderTarget=1)
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-10-26 05:20:47.050658932 +0000 UTC m=+238.651302118
        ScaleDown:   NoCandidates (candidates=0)
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-10-26 05:20:47.050658932 +0000 UTC m=+238.651302118

        Name:        obscured
        Health:      Healthy (ready=1 unready=0 notStarted=0 longNotStarted=0 registered=1 longUnregistered=0 cloudProviderTarget=1 (minSize=1, maxSize=100))
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-10-26 05:20:47.050658932 +0000 UTC m=+238.651302118
        ScaleUp:     NoActivity (ready=1 cloudProviderTarget=1)
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-10-26 05:20:47.050658932 +0000 UTC m=+238.651302118
        ScaleDown:   NoCandidates (candidates=0)
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-10-26 05:20:47.050658932 +0000 UTC m=+238.651302118

        Name:        obscured
        Health:      Healthy (ready=1 unready=0 notStarted=0 longNotStarted=0 registered=1 longUnregistered=0 cloudProviderTarget=1 (minSize=1, maxSize=100))
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-10-26 05:20:47.050658932 +0000 UTC m=+238.651302118
        ScaleUp:     NoActivity (ready=1 cloudProviderTarget=1)
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-11-09 08:49:17.188499976 +0000 UTC m=+1222348.789143150
        ScaleDown:   NoCandidates (candidates=0)
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-11-09 08:53:30.879288373 +0000 UTC m=+1222602.479931549

        Name:        obscured
        Health:      Healthy (ready=1 unready=0 notStarted=0 longNotStarted=0 registered=1 longUnregistered=0 cloudProviderTarget=1 (minSize=1, maxSize=100))
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-10-26 05:20:47.050658932 +0000 UTC m=+238.651302118
        ScaleUp:     NoActivity (ready=1 cloudProviderTarget=1)
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-11-07 00:10:41.430022289 +0000 UTC m=+1018433.030665467
        ScaleDown:   NoCandidates (candidates=0)
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-11-07 00:12:53.55169393 +0000 UTC m=+1018565.152337097

        Name:        obscure
        Health:      Healthy (ready=1 unready=0 notStarted=0 longNotStarted=0 registered=1 longUnregistered=0 cloudProviderTarget=1 (minSize=1, maxSize=100))
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-10-26 05:20:47.050658932 +0000 UTC m=+238.651302118
        ScaleUp:     NoActivity (ready=1 cloudProviderTarget=1)
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-11-08 08:48:47.137796002 +0000 UTC m=+1135918.738439166
        ScaleDown:   NoCandidates (candidates=0)
                     LastProbeTime:      2021-11-10 00:01:27.486224515 +0000 UTC m=+1277079.086867680
                     LastTransitionTime: 2021-11-08 08:49:36.858799215 +0000 UTC m=+1135968.459442400

  kind: ConfigMap
  metadata:
    annotations:
      cluster-autoscaler.kubernetes.io/last-updated: 2021-11-10 00:01:27.986516659
        +0000 UTC
    creationTimestamp: "2021-05-04T15:14:36Z"
    managedFields:
    - apiVersion: v1
      fieldsType: FieldsV1
      fieldsV1:
        f:data:
          .: {}
          f:status: {}
        f:metadata:
          f:annotations:
            .: {}
            f:cluster-autoscaler.kubernetes.io/last-updated: {}
      manager: cluster-autoscaler
      operation: Update
      time: "2021-05-04T15:14:36Z"
    name: cluster-autoscaler-status
    namespace: kube-system
- apiVersion: v1
  data:
    root-cert.pem: |
      -----BEGIN CERTIFICATE-----
      MIIC/DCCAeSgAwIBAgIQYfZd7aU845ZUCPqVw6YoyDANBgkqhkiG9w0BAQsFADAY
      MRYwFAYDVQQKEw1jbHVzdGVyLmxvY2FsMB4XDTIxMTExMjAxMzkzOFoXDTMxMTEx
      MDAxMzkzOFowGDEWMBQGA1UEChMNY2x1c3Rlci5sb2NhbDCCASIwDQYJKoZIhvcN
      AQEBBQADggEPADCCAQoCggEBAMcJzGjz5YNyf2r/O683FjNsFtxGnQgKJVxIpakq
      gJziiJh6mfAI+piqfkJUDZoq2tGwnFHEpLil7Vb016ubdez7Vd3lbMYyY8Gn0JU2
      RDaP+NNhj4KwQv2QN+ebz3NVwdnJUog7aD7E7cI+CBHJ2q3z+UQCSQ2NjZPTqxI5
      xl9AIz5MQN9kjMG6oooexmKNqzW93CxQfceNrZWehibSzKQCLrKLhPa2OAMP6aqt
      YINBLz5/y3QvN9qectAWCEgU6+HOtCJKAc4h3ZjZwLjq5yxOPJ0K2ceNx/sGk5ya
      bc4gIROeiIfR8QauaEeDzwDmtPoX8YnMdRWtiEfpUS5nqdsCAwEAAaNCMEAwDgYD
      VR0PAQH/BAQDAgIEMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFL9uVSgnggVI
      +qdB/E/bMehwNerLMA0GCSqGSIb3DQEBCwUAA4IBAQC9OVksSsuVMaR5azJfrK3e
      L5Kl9Rx1T5+VxZ2fxxCmf3wyPmc9VjDnsb+FtmjQDs7IAt7/hxCHys7mC8ZBCLbf
      dOu91TuHzGmGP0NBOmWBFgeRSHnIwYB+mKAtTtlOpZwGdJhjvKcC2eSR2oou6I9j
      SZiVkv6FcEvGdzcYtqw5mulX9jM+zLMl76Bsm4I8ZvhmzUPwmUqaWpz3tQ88IrHc
      xOZSGJ13nTmIgaWrzRS5mxerY4wRSTOLPC8FvWZVToVoM8ZuMfM5pMhalb764E6q
      78UTleJ3XpFojiRqsFc20WuHAjtfwgz4vktmXo8Hck5UJs/f1aKfKIZgPAVSXaVj
      -----END CERTIFICATE-----
  kind: ConfigMap
  metadata:
    creationTimestamp: "2021-11-12T01:39:39Z"
    labels:
      istio.io/config: "true"
    name: istio-ca-root-cert
    namespace: default
    resourceVersion: "5288"
    uid: 992bb94c-c1c9-404e-bdec-87e62b853b79
- apiVersion: v1
  data:
    ca.crt: |
      -----BEGIN CERTIFICATE-----
      MIIC/jCCAeagAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl
      cm5ldGVzMB4XDTIxMTExMjAwNTQyNVoXDTMxMTExMDAwNTQyNVowFTETMBEGA1UE
      AxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL5H
      /7DMS34Un4AB/b1Ykh8EzIfc3vawaz0vcSXhxBazrwLaVBp1CjOZf4zPzWOEQ6a5
      s3gYMa16n05dN1LJNKBxOvmEetuD54pqwpXVXuCN+wt2S3DsssKOresYAcJPf7Qf
      YLVrQ5ZPftj5utN+0xcaSl7YC+bAtVnu/i+ZXPrGS/0uGOAL2gCZs1xcazmOzSwK
      HvR78b4vu8ch29huehIwT8XPzfH5EZyd2jgaLwhQUlzm182z1IiwGM9YGXhlSAtn
      ksrlwctf+ARbYKV4+OS9YZ36qLoqktQemMur18t67iuWiomeoaUJ3skPijotl5Kk
      dwSdFKn5OWilnI+104UCAwEAAaNZMFcwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB
      /wQFMAMBAf8wHQYDVR0OBBYEFKocHx1yer/illx/cr4yJ2afUExQMBUGA1UdEQQO
      MAyCCmt1YmVybmV0ZXMwDQYJKoZIhvcNAQELBQADggEBABLoNrCvqHb32PoZra4Z
      qznXiIJd4s5WfQudtLypdnXX9NvE/uOASLCxc3hjpTzii6wiKLRThLq+CaznNwi0
      gVNkPielMhEmkD8x/NxSKRlOr94pwxmFGNWNpjkkuLaxQRUlLI1qDa/tio4Fo7YE
      RahAw0Aa1Mwh/+49xPFpFk+RYOEBssTOOzBi69mhaGbfGcA3YtLZ3R6liqmVqKhD
      ovtWj0CWyJUB7rLKKhwzYvv6MrYZ87WYeT8zTHq8L8EhbdsDOvuFs/IK2BbQdyd2
      7aXsHT08dc1QCH3rP2bdkehcLtGXSP++7gJNXxL280EnyWFt+KFkikbLLI/s+oVE
      1gc=
      -----END CERTIFICATE-----
  kind: ConfigMap
  metadata:
    annotations:
      kubernetes.io/description: Contains a CA bundle that can be used to verify the
        kube-apiserver when using internal endpoints such as the internal service
        IP or kubernetes.default.svc. No other usage is guaranteed across distributions
        of Kubernetes clusters.
    creationTimestamp: "2021-11-12T00:54:57Z"
    name: kube-root-ca.crt
    namespace: default
    resourceVersion: "422"
    uid: 193862c7-8b99-4838-889e-9e1e3a15c930
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""
`,
			want: []string{
				"ConfigMap/cluster-autoscaler-status.kube-system",
				"ConfigMap/istio-ca-root-cert.default",
				"ConfigMap/kube-root-ca.crt.default",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := bytes.Buffer{}
			if err := GrepResources(Selector{}, strings.NewReader(tt.input), &o, Summary, false, false); err != nil {
				t.Fatal(err)
			}
			l := strings.Split(o.String(), "\n")
			if l[len(l)-1] == "" {
				l = l[:len(l)-1]
			}
			if !reflect.DeepEqual(l, tt.want) {
				t.Errorf("got = %v, want %v", l, tt.want)
			}
		})
	}
}
