console.log("Hey");

export function sendMessage() {
    let text = document.getElementById("input").value
    if (text != "" && !checkEmptyMessage(text)) {
        let forum = document.getElementById("forum")
        let element = document.createElement("div")
        element.innerText += text
        forum.appendChild(element)
        document.getElementById("input").value = ""
    } else {
        document.getElementById("input").value = ""
        console.log("no text")
    }
}

function checkEmptyMessage(message) {
    return !/\S/.test(message)
}