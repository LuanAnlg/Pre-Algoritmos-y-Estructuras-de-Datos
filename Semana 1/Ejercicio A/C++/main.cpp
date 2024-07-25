#include <iostream>

int main() {
    while (true) {
        int n;
        std::cout << "Introduce el número de personas (0 para salir):" << std::endl;
        std::cin >> n;

        if (n == 0) {
            return 0;
        }

        if (n < 1 || n >= 100) {
            std::cout << "Error: entrada inválida. Por favor, introduce un número válido." << std::endl;
            continue;
        }

        int edades[n];

        std::cout << "Introduce " << n << " edades separadas por espacio:" << std::endl;
        bool valid = true;

        for (int i = 0; i < n; ++i) {
            if (!(std::cin >> edades[i])) {
                std::cout << "Error: entrada de edad inválida. Por favor, introduce los valores nuevamente." << std::endl;
                valid = false;
                break;
            }
        }

        if (!valid) {
            continue;
        }

        for (int i = 0; i < n - 1; ++i) {
            for (int j = 0; j < n - 1 - i; ++j) {
                if (edades[j] > edades[j + 1]) {
                    std::swap(edades[j], edades[j + 1]);
                }
            }
        }

        for (int i = 0; i < n; ++i) {
            if (i > 0) {
                std::cout << " ";
            }
            std::cout << edades[i];
        }
        std::cout << std::endl;
    }
}