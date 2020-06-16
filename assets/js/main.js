document.getElementById("boring-terminal-text").focus();

if (document.activeElement) {
    document.addEventListener("keypress", function (event) {
        if (event.keyCode === 13) {
            alert("Enter is pressed")
            event.preventDefault()
        }
    })
}