//Création des éléments, attribution des classes, id, et du texte contenu le cas échéant

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

const submitButton = document.createElement("button")
submitButton.classList.add("repondre")
submitButton.innerText = "Répondre"

const rightPart = document.createElement("right")
rightPart.classList.add("right")

const title = document.createElement("div")
title.classList.add("title-in")
title.innerText = element.Title

const description = document.createElement("div")
description.classList.add("text-in")
description.innerText = element.Contain


//Ajout des éléments les uns dans les autres

leftPart.appendChild(userNameContainer)
leftPart.appendChild(br)
leftPart.appendChild(br)
leftPart.appendChild(hr)
leftPart.appendChild(submitButton)

rightPart.appendChild(title)
rightPart.appendChild(hr)
rightPart.appendChild(description)

card.appendChild(leftPart)
card.appendChild(rightPart)

divContent.appendChild(card)

mainContainer.appendChild(divContent)





mainContainer.appendChild(divContent)