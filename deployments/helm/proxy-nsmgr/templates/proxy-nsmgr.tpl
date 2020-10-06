---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: proxy-nsmgr
  namespace: {{ .Release.Namespace }}
spec:
  selector:
    matchLabels:
      app: proxy-nsmgr-daemonset
  template:
    metadata:
      labels:
        app: proxy-nsmgr-daemonset
    spec:
      serviceAccount: proxy-nsmgr-acc
      containers:
        - name: proxy-nsmd
          image: {{ .Values.registry }}/{{ .Values.org }}/proxy-nsmd:{{ .Values.tag }}
          imagePullPolicy: {{ .Values.pullPolicy }}
          ports:
            - containerPort: 5006
              hostPort: 5006
          env:
            - name: INSECURE
              value: {{ .Values.insecure | default false | quote }}
            - name: PROXY_NSMD_K8S_REMOTE_PORT
              value: {{ .Values.remoteNsrPort | quote }}
          volumeMounts:
            - name: spire-agent-socket
              mountPath: /run/spire/sockets
              readOnly: true
        - name: proxy-nsmd-k8s
          image: {{ .Values.registry }}/{{ .Values.org }}/proxy-nsmd-k8s:{{ .Values.tag }}
          imagePullPolicy: {{ .Values.pullPolicy }}
          ports:
            - containerPort: 5005
              hostPort: 5005
          env:
            - name: NODE_NAME
              valueFrom:
                fieldRef:
                  fieldPath: spec.nodeName
            - name: INSECURE
              value: {{ .Values.insecure | default false | quote }}
            - name: TRACER_ENABLED
              value: {{ .Values.global.JaegerTracing | default false | quote }}
            - name: JAEGER_AGENT_HOST
              value: jaeger.nsm-system
            - name: JAEGER_AGENT_PORT
              value: "6831"
            - name: PROXY_NSMD_K8S_REMOTE_PORT
              value: {{ .Values.remoteNsrPort | quote }}
          volumeMounts:
            - name: spire-agent-socket
              mountPath: /run/spire/sockets
              readOnly: true
      volumes:
        - hostPath:
            path: /run/spire/sockets
            type: DirectoryOrCreate
          name: spire-agent-socket
---
apiVersion: v1
kind: Service
metadata:
  name: pnsmgr-svc
  labels:
    app: proxy-nsmgr-daemonset
  namespace: {{ .Release.Namespace }}
spec:
  type: NodePort
  ports:
    - name: pnsmr
      port: 5005
      nodePort: 31505
      protocol: TCP
    - name: pnsmd
      port: 5006
      nodePort: 31506
      protocol: TCP
  selector:
    app: proxy-nsmgr-daemonset
