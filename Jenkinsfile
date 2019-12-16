node {
    def app

    stage('Clone repository') {
        /* repository cloned to our workspace */
        checkout scm
    }

    stage('Build image') {
        /* builds the docker image */
        docker.withRegistry('https://spinreg-docker-registry.spinnaker:5000')    {
            app = docker.build("hostinfo:${env.BUILD_ID}")
    }

    stage('Push image') {
        /* push the image */
        docker.withRegistry('https://spinreg-docker-registry.spinnaker:5000')    {
            app.push()
        }
    }
}
