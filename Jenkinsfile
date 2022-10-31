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
		}
		stage('Deploy') {
		}
	}
}
