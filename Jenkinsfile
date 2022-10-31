pipeline {
    agent any
	stages {
		stage('Build') {
			sh 'cd ~/workspace'
			sh 'helm create daemonset'	
		}
		stage('Docker Upload') {
		}
		stage('Deploy') {
		}
	}
}
