MVN_TEST=mvn test -Dtags="!unimplemented"
NODE=1

test:
	$(MVN_TEST)

test-para:
	$(MVN_TEST) -DinParallel -Dnodes=$(NODE)

clean:
	mvn clean

test-clean: clean
	$(MVN_TEST)

test-compile:
	mvn test-compile

