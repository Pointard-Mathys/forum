const URL = "/api-reponses"
// let topicID = localStorage.getItem('Data');

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
        // console.log("Page chat : ", topicID);
        console.log("element.Id_topic : ", element.Id_topic, element.TextContent);
        const queryString = window.location.search;
        const urlParams = new URLSearchParams(queryString);
        const topicId = urlParams.get("topicId")
        console.log("topicId : ", topicId);
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