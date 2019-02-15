@Library('stash@loy/loyalty-pipeline-library') _

import com.zooplus.zpp.jplib.loy.*
import com.zooplus.zpp.jplib.loy.pipeline.*

deploymentDefinition = new DeploymentDefinition(
        pdc: "sam",
        productGroup: "loyalty",
        productName: "common",
        applicationName: "lmwtfy",
        baseFolder: "lmwtfy/",
        deploymentParams: [
                PROD: [
                        APP_PORT     : 5222,
                        APP_INSTANCES: 1,
                        APP_MEMORY   : 512
                ],
                DEV : [
                        APP_PORT     : 5222,
                        APP_INSTANCES: 1
                ]
        ]
)

loy.pipeline.builder()
        .withLabel('golang')
        .withCheckout()
        .withDockerPublication(deploymentDefinition)
        .withMarathonProdDeployment(deploymentDefinition, [])
        .run()