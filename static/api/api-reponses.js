fetch("/api-reponses",
{
    method: "GET",
    headers: {
        "content-type": "application/json"
    },
    
}).then((res) => {
    return res.json()
})
.then((res) => {
    const abs = document.getElementById("absolute")

    abs.innerText = res.test
})