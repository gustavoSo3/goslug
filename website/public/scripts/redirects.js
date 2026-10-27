


const subBtn = document.getElementById("submit-new-redirect");

const apiRedirectURL = "/api/redirects";

subBtn.addEventListener("click", async () => {
    await postLogIn();
});

document.addEventListener("keydown", async (event) => {
    if(event.key === "Enter"){
        await postLogIn();
    }
});


async function postLogIn(){

    const newRedirectDescription = document.getElementById("new-description");
    const newRedirectSlug = document.getElementById("new-slug");
    const newRedirectDestination = document.getElementById("new-destination");

    if( newRedirectDescription.value === "" || newRedirectSlug.value === "" || newRedirectDestination.value === ""){
        return
    }
    const newRedirect = {
        Description: newRedirectDescription.value,
        Slug: "/" + newRedirectSlug.value,
        Destination: newRedirectDestination.value,
    };

    console.log(newRedirect);

    try {
        // Make the POST request to /api/redirects
        const res = await fetch(apiRedirectURL, {
        method: 'POST', // Specify the HTTP method
        headers: {
            'Content-Type': 'application/json' // Tell the server you're sending JSON
        },
        body: JSON.stringify(newRedirect) // Convert the JS object into a JSON string
        });

        // Handle the server's response
        if (res.ok) {
            const newConfirmedRedirect = await res.json();
            console.log('Success:', newConfirmedRedirect);

            newRedirectDescription.value = "";
            newRedirectSlug.value = "";
            newRedirectDestination.value = "";

            location.reload();
            

        } else {
            console.error('Server error:', res.status);
        }
    } catch (error) {
        // Handle network errors or issues with the fetch itself
        console.error('Network error:', error);
    }
};

async function deleteRedirect(slug) {

    const deleteSlug = slug

    console.log(deleteSlug);

    if (confirm(`Are you shure you want to delete the slug: "${slug}"`)){
        try {
            // Make the DELETE request to /api/redirects
            const res = await fetch(apiRedirectURL, {
            method: 'DELETE', // Specify the HTTP method
            headers: {
                'Content-Type': 'application/json' // Tell the server you're sending JSON
            },
            body: JSON.stringify(deleteSlug) // Convert the JS object into a JSON string
            });

            // Handle the server's response
            if (res.ok) {
                const newConfirmedRedirect = await res.body;
                console.log('Success:', newConfirmedRedirect);
                location.reload();

                
            } else {
                console.error('Server error:', res.status);
            }
        } catch (error) {
            // Handle network errors or issues with the fetch itself
            console.error('Network error:', error);
        }
    }
};