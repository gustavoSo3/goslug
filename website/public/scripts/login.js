const subBtn = document.getElementById("submit-login");

const loginURL = "/login";

subBtn.addEventListener("click", async () => {
    await postLogIn();
});

document.addEventListener("keydown", async (event) => {
    if(event.key === "Enter"){
        await postLogIn();
    }
});


async function postLogIn() {
    
    const username = document.getElementById("user-name");
    const password = document.getElementById("password");


    const login = {
        username: username.value,
        password: password.value,
    };

    console.log(login);

    try {
        // Make the POST request to /api/redirects
        const res = await fetch(loginURL, {
        method: 'POST', // Specify the HTTP method
        headers: {
            'Content-Type': 'application/json' // Tell the server you're sending JSON
        },
        body: JSON.stringify(login) // Convert the JS object into a JSON string
        });

        // Handle the server's response
        if (res.ok) {

            location.href = "/admin"
            console.log(document.cookie)

        } else {
            console.error('Server error:', res.status);
            username.value = "";
            password.value = "";
            location.reload()
        }
    } catch (error) {
        // Handle network errors or issues with the fetch itself
        console.error('Network error:', error);
    }
}