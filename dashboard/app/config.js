// This is the default configuration for the dev mode of the web dashboard.
// A generated version of this config is created at run-time when running
// the web dashboard from the binary
window.STRAYARD_CONFIG = {
  api: {
	strayard: {
	  hostPort: "localhost:8848",
	  prefix: "/syapi"
	},
    openshift: {
      hostPort: "localhost:8848",
      prefix: "/osapi"
    },
    k8s: {
      hostPort: "localhost:8848",
      prefix: "/api"
    }
  },
  auth: {
  	oauth_authorize_uri: "https://localhost:8848/oauth/authorize",
  	oauth_redirect_base: "https://localhost:9000",
  	oauth_client_id: "strayard-web-dashboard",
  	logout_uri: ""
  },
};