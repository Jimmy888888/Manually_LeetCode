CXX = g++
CXXFLAGS = -std=c++11 -Wall -Wextra

solution: solution.cpp
	$(CXX) $(CXXFLAGS) -o solution solution.cpp

test: test_solution.cpp
	$(CXX) $(CXXFLAGS) -o test_solution test_solution.cpp
	./test_solution

clean:
	rm -f solution test_solution

.PHONY: test clean