#include <unordered_map>
#include <set>
#include <string>
#include <vector>
#include <iostream>
#include <limits>
#include <algorithm>

std::set<int> countPrimes(int n) {
    using namespace std;
    if (n <= 2) {
        return set<int>();
    }
    vector<int> is_prime(n, true);

    for (int i = 3; i * i < n; i += 2) {
        if (!is_prime[i]) {
            continue;
        }

        for (int j = i * i; j < n; j += 2 * i) {
            if (!is_prime[j]) {
                continue;
            }
            is_prime[j] = false;
        }
    }
    set<int> result;
    result.insert(2);
    for (int i = 3; i <= n; ++++i) {
        if (is_prime[i] == true)
            result.insert(i);
    }
    return result;
}

std::string count(const int& length, const std::set<int>& primeSet) {
    std::vector<int> input(length);
    std::vector<int> result(length + 1);
    std::set<int> dcryptKey;

    for (int i = 0; i < length; ++i) {
        std::cin >> input[i];
    }

    //auto minCryptoText = std::numeric_limits<int>::max();
    {
        auto iter = primeSet.begin();
        for (; iter != primeSet.end(); iter++) {
            if (input[0] % *iter == 0)
                break;
        }

        auto startIdx = 0;
        for (auto i = 0; i < input.size(); ++i, ++startIdx) {
            if (input[i] != input[i + 1])
                break;
        }
        int left = *iter, right = input[startIdx] / *iter;
        if (input[startIdx + 1] % left == 0)
            std::swap(left, right);
        result[startIdx] = left;
        //minCryptoText = left;
        dcryptKey.insert(left);

        for (auto i = startIdx + 1; i < length; ++i) {
            left = right;
            right = input[i] / left;
            result[i] = left;
            dcryptKey.insert(left);
            //minCryptoText = std::min(minCryptoText, result[i]);
        }
        result[length] = right;
        dcryptKey.insert(right);
        //minCryptoText = std::min(minCryptoText, result[length]);

        if (startIdx != 0) {
            for (auto i = startIdx - 1; i >= 0; --i) {
                result[i] = input[i] / result[i + 1];
                dcryptKey.insert(result[i]);
                //minCryptoText = std::min(minCryptoText, result[i]);
            }
        }
    }

    std::unordered_map<int, char> decryptMap;
    {
        {
            int i = 0;
            for (const auto& v : dcryptKey) {
                decryptMap[v] = 'A' + i;
                ++i;
            }
        }

        std::string decryptText;

        for (const auto& v : result) {
            decryptText.push_back(decryptMap[v]);
        }
        return decryptText;
    }
}

int main() {
    using namespace std;

    static const auto __ = []()
    {
        ios::sync_with_stdio(false);
        cin.tie(0);
        return 0;
    }();

    int testTime;
    cin >> testTime;
    const auto primeSet = countPrimes(10000);
    for (int i = 0; i < testTime; ++i) {
        int N;
        cin >> N; //useless
        cin >> N;
        const auto result = count(N, primeSet);
        cout << "Case #" << i + 1 << ": " << result << endl;
    }
    return 0;
}
