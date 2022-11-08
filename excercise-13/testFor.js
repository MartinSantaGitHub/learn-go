const testFor = (n) => {
    for (let i = 0; i < n; i++) {
        //console.log(i);
    }
};

const init = new Date();

testFor(10000000);

const end = new Date();
const delta = end - init;

console.log(delta);
