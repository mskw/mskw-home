pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'echo "Hello World"'
                sh '''
                    kill `cat /www/pid/mskw-home`
                    echo "Multiline shell steps works too"
                    cat README.md
                    ls -lah
                    go run src/main.go &
                '''
            }
        }
    }
}