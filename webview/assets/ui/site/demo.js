const output = document.getElementById("count");
const hint = document.getElementById("hint");

let count = 0;

// One function serves the page's button and Go's WebView.Eval.
function bump() {
  count += 1;
  output.textContent = String(count);
  goiPost(`ping ${count}`);
  hint.textContent = `Go was told: ping ${count}`;
}

document.getElementById("ping").addEventListener("click", bump);
window.demoBump = bump;

// Go's status line shows this, so a load can be seen from the other side.
goiPost("ready");
