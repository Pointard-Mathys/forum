const URL = "/test2"

export let tmp 
export let topicID
// export const responseMessages = topic => {
//     console.log("Topic :", topic);
//     return topic
// }

export function sendMessage() {
    getApi(URL)
    let text = document.getElementById("message").value
    let forum = document.getElementById("forum")
    if (text != "" && !checkEmptyMessage(text)) {
        const element = document.createElement("div")
        element.innerText += text
        forum.appendChild(element)
    } else {
        document.getElementById("message").value = ""
        alert("You're trying to send an empty message!")
    }
}

function checkEmptyMessage(message) {
    return !/\S/.test(message)
}

export function getApi(URL) {
    
    fetch(URL)
    .then((response) => response.json())
    .then((responseMessages) => {
        tmp = responseMessages
        printAPI(tmp)
    })
}

export function printAPI(topic) {
    console.log("Topic : ", topic[0].User_Name);
    topic.forEach(element => {
        //Création des éléments, attribution des classes, id, et du texte contenu le cas échéant
        topicID = element.Id
        console.log("Page addmessage : ", topicID);

        const mainContainer = document.getElementById("topic-container")
        
        const divContent = document.createElement("div")
        divContent.classList.add("content")
        
        const card = document.createElement("div")
        card.classList.add("card")
        
        const leftPart = document.createElement("div")
        leftPart.classList.add("left")
        leftPart.setAttribute("id", "left")
        
        const userNameContainer = document.createElement("h1")
        userNameContainer.classList.add("pseudo")
        userNameContainer.innerText = element.User_Name
        
        const br = document.createElement("br")
        const hr = document.createElement("hr")

        const a = document.createElement("a")
        a.setAttribute("href", "/messages")
        
        const submitButton = document.createElement("button")
        submitButton.classList.add("repondre")
        submitButton.setAttribute("name", "idTopic")
        submitButton.setAttribute("value", element.Id)
        submitButton.innerText = "Répondre"
        console.log("ID : ", element.Id, element.Contain);
        // submitButton.addEventListener("click", localStorage.setItem('Data', element.Id))
        
        const rightPart = document.createElement("right")
        rightPart.classList.add("right")
        
        const title = document.createElement("div")
        title.classList.add("title-in")
        title.innerText = element.Titre
        
        const description = document.createElement("div")
        description.classList.add("text-in")
        description.innerText = element.Contain
        
        
        //Ajout des éléments les uns dans les autres

        a.appendChild(submitButton)
        
        leftPart.appendChild(userNameContainer)
        leftPart.appendChild(br)
        leftPart.appendChild(br)
        leftPart.appendChild(hr)
        leftPart.appendChild(a)
        
        rightPart.appendChild(title)
        rightPart.appendChild(hr)
        rightPart.appendChild(description)
        
        card.appendChild(leftPart)
        card.appendChild(rightPart)
        
        divContent.appendChild(card)
        
        mainContainer.appendChild(divContent)
    });
}

getApi(URL)
