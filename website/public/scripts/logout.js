const logoutBtn = document.getElementById("logout-button");

const loginURL = "/login";

logoutBtn.addEventListener("click", async () => {

    // Make the POST request to /api/redirects
    const res = await fetch(loginURL, {
    method: 'DELETE',
    });

    // Handle the server's response
    if (res.ok) {

        location.href = "/"

    } else {
        location.reload()
    }
});