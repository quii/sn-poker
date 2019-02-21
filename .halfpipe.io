team: engineering-enablement
pipeline: sn-poker
#slack_channel: '#poker'

feature_toggles:
- update-pipeline

tasks:

- type: docker-compose
  name: build
  save_artifacts: [.]
  
- type: deploy-cf
  name: deploy
  api: ((cloudfoundry.api-snpaas))
  space: test
  deploy_artifact: cmd/poker-app
  vars:
    BIN: ((sn-poker.jsbin))
