const fibonacci = (n) => {
    if (n > 1) {
        return fibonacci(n - 2) + fibonacci(n - 1);
    }

    return n;
};

const init = new Date();

fibonacci(42);

const end = new Date();
const delta = end - init;

console.log(delta);
