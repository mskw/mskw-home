pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'echo "Hello World"'
                sh '''
                    if [ ! -f "/www/pid/mskw-home" ]; then
                        kill `cat /www/pid/mskw-home`
                    fi
                    echo "Multiline shell steps works too"
                    cat README.md
                    ls -lah
                    go run src/main.go &
                '''
            }
        }
    }
}