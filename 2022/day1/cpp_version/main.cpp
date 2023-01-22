#include <fstream>
#include <iostream>
#include <queue>
#include <string>

int main(int argc, char *argv[]) {
  if (argc == 1)
    return 1;
  std::priority_queue<int> heap;
  std::fstream file(argv[1]);

  std::string buffer;
  int sum = 0;
  while (std::getline(file, buffer)) {
      //std::cout << buffer << '\n';
    if (buffer.empty()) {
      heap.push(sum);
      sum = 0;
    } else {
      sum += std::stoi(buffer);
    }
  }

  int i = 0;
  int result = 0;
  while(heap.size() != 0 && i < 3) {
      result += heap.top();
      heap.pop();
      ++i;
  }
  std::cout << result << '\n';

  return 0;
}
