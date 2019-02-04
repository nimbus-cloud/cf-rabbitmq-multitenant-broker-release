###### Help ###################################################################

help:
	@echo '    deps ...................................... installs dependencies'
	@echo '    go-vet .................................... runs go vet in source code'
	@echo '    integration_tests ......................... runs integration tests'
	@echo '    unit_tests ................................ runs unit tests'
	@echo '    broker   .................................. builds the broker binary'

###### Dependencies ###########################################################

deps:
			git submodule update --init --recursive

###### Tools ###############################################################

go-vet:
				go vet `go list ./... | grep -v vendor`

###### Tests ###############################################################

integration_tests:
	GO111MODULE=on GOFLAGS='-mod=vendor' ginkgo -randomizeSuites=true -randomizeAllSpecs=true -keepGoing=true -r -race -p integrationtests/

unit_tests:
	GO111MODULE=on GOFLAGS='-mod=vendor' ginkgo -randomizeSuites=true -randomizeAllSpecs=true -keepGoing=true -r -race -p --skipPackage=integrationtests .

############################################################################

broker:
			go build -o cmd/service-broker/service-broker ./cmd/service-broker
