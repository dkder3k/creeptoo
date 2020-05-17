document.getElementById("encrypt").addEventListener("click", encrypt);
document.getElementById("decrypt").addEventListener("click", decrypt);

function getCipher() {
    const radios = document.getElementsByName("cipher");
    for (let i = 0; i < radios.length; i++) {
        if (radios[i].checked) return radios[i].value;
    }
}

function transform(action) {
    const inputText = document.getElementById("input").value;
    const key = document.getElementById("key").value;
    const cipher = getCipher();

    const Http = new XMLHttpRequest();
    const url = `http://crpt.formyown.xyz/api/v1/${cipher}?key=${key}&text=${inputText}&action=${action}`;
    Http.open("GET", url);
    Http.send();

    Http.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            document.getElementById("result").innerHTML = Http.responseText;
        }
    }
}

function decrypt() {
    transform("decrypt");
}

function encrypt() {
    transform("encrypt");
}
