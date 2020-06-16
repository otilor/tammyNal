document.getElementById("boring-terminal-text").focus();

if (document.activeElement) {
    document.addEventListener("keypress", function (event) {
        if (event.keyCode === 13) {
            let log = document.getElementById("answer")
            log.innerHTML = "Hey";
            event.preventDefault()
        }
    })
}