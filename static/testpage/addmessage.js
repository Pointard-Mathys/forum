console.log("Hey");

export function sendMessage() {
    let text = document.getElementById("message").value
    let forum = document.getElementById("forum")
    if (text != "" && !checkEmptyMessage(text)) {
        const element = document.createElement("div")
        element.innerText += text
        forum.appendChild(element)
    } else {
        document.getElementById("message").value = ""
        console.log("hya");
        alert("You're trying to send an empty message!")
    }
}

function checkEmptyMessage(message) {
    return !/\S/.test(message)
}