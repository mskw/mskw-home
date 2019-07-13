pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'echo "Hello World"'
                sh '''
                    if [ -f "/www/pid/mskw-home" ]; then
                        kill `cat /www/pid/mskw-home`
                    fi
                    cat README.md

                '''
            }
        }
        stage('Deploy') {
            steps {
                withEnv(['JENKINS_NODE_COOKIE=dontkillme']) {
                    sh 'go run src/main.go'
                }
            }
        }
    }
}