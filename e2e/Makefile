MVN_TEST=mvn test -Dtags="!unimplemented"
NODE=1

test:
	$(MVN_TEST) -Dflags='--max-retries-count=3'

test-para:
	$(MVN_TEST) -DinParallel -Dnodes=$(NODE)

clean:
	mvn clean

test-clean: clean
	$(MVN_TEST)

test-compile:
	mvn test-compile

