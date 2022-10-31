pipeline {
    agent any
	stages {
		stage('Build') {
			steps {
				sh 'cd ~/workspace'
				sh 'helm create daemonset'	
			}
		}
		stage('Docker Upload') {
			steps {
				sh 'echo Docker Upload'
			}
		}
		stage('Deploy') {
			steps {
				sh 'echo Deploy'
			}
		}
	}
}
