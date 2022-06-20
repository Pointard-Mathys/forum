const URL = "/api-reponses"

export function getApi(URL) {
    let tmp

    fetch(URL)
        .then((response) => response.json())
        .then((responseMessages) => {
            tmp = responseMessages
            printAPI(tmp)
        })
}

function printAPI(reponse) {
    const messageList = document.getElementById("chat-thread")
    reponse.forEach(element => {
        const queryString = window.location.search;
        const urlParams = new URLSearchParams(queryString);
        const topicId = urlParams.get("topicId")
        if (element.Id_topic == topicId) {
            const pseudoContainer = document.createElement("div")
            pseudoContainer.classList.add("pseudo")
            pseudoContainer.innerText = element.User_name

            const messageContainer = document.createElement("li")
            messageContainer.innerText = element.TextContent

            messageList.appendChild(pseudoContainer)
            messageList.appendChild(messageContainer)
        }
    });

}

getApi(URL)