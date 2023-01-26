import { readFileSync } from "fs";
import Heap from "heap";

function main() {
    const heap = new Heap(function(a: number, b: number) {
        return b - a;
    })
    const content = readFileSync("./content.txt").toString();
    let sum = 0;
    content.split('\n').forEach(line => {
        if(line.length == 0) {
            heap.push(sum);
            sum = 0;
        } else {
            sum += parseInt(line);
        }
    });

    let i = 0;
    let result = 0;
    while(heap.size() > 0 && i < 3) {
        const value = heap.pop();
        if(value == undefined) break;
        result += value;
        ++i;
    }
    console.log(result)
}

main();
