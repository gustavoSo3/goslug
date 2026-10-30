


const newRedirectBtn = document.getElementById("submit-new-redirect");
const liveAlertDiv = document.getElementById('redirects-live-alert')

const newRedirectDescription = document.getElementById("new-description");
const newRedirectSlug = document.getElementById("new-slug");
const newRedirectDestination = document.getElementById("new-destination");

const apiRedirectURL = "/api/redirects";

if (newRedirectBtn) {
    newRedirectBtn.addEventListener("click", async () => {
        if (AreValidFields(newRedirectDescription.value, newRedirectSlug.value, newRedirectDestination.value)) {
            await postRedirect(newRedirectDescription.value, newRedirectSlug.value, newRedirectDestination.value);
        }
    });
}

document.addEventListener("keydown", async (event) => {
    if (AreValidFields(newRedirectDescription.value, newRedirectSlug.value, newRedirectDestination.value)) {
        if(event.key === "Enter"){
            await postRedirect(newRedirectDescription.value, newRedirectSlug.value, newRedirectDestination.value);
        }
    }
});

document.addEventListener("keyup", () => {
    if (AreValidFields(newRedirectDescription.value, newRedirectSlug.value, newRedirectDestination.value)) {
        newRedirectBtn.disabled = false;
    } else {
        newRedirectBtn.disabled = true;
    }
});



function AreValidFields(description, slug, destination) {

    description = description.trim();
    slug = slug.trim();
    destination = destination.trim();

    // If any empty string
    if (description === "" || slug === "" || destination === "") return false;


	// Description is valid if is less than 100 chars
    if (description.length > 100) return false;

	// Slug is valid if starts with "/" is less than 20 chars and doesnt have any spaces (maybe will need to check for -/_() later)
    if (slug.length > 20 || slug.includes(" ")) return false;

	// Must be a valid URL
    try {
        const u = new URL(destination);
        return u.protocol === "http:" || u.protocol === "https:";
    } catch (err) {
        return false;
    }

	return true;
}

async function postRedirect(description, slug, destination){

    description = description.trim();
    slug = slug.trim();
    destination = destination.trim();

    const newRedirect = {
        Description: description,
        Slug: "/" + slug,
        Destination: destination,
    };

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
            location.reload();
        }else{
            handleBadRedirectStatus(res.status);
        }
    } catch (error) {
        appendAlert("Server error: " + error , "danger");
    }
};

async function deleteRedirect(slug) {

    if (confirm(`Are you sure you want to delete the Slug: "${slug}"`)){
        try {
            // Make the DELETE request to /api/redirects
            const res = await fetch(apiRedirectURL, {
            method: 'DELETE', // Specify the HTTP method
            headers: {
                'Content-Type': 'application/json' // Tell the server you're sending JSON
            },
            body: JSON.stringify(slug) // Convert the JS object into a JSON string
            });

           
            // Handle the server's response
            if (res.ok) {
                location.reload();
            }else {
                handleBadRedirectStatus(res.status);
            }
        }catch (error) {
            appendAlert("Server error: " + error , "danger");
        }
    }
};

//redirects-live-alert
function appendAlert(message, type)  {
  const wrapper = document.createElement('div')
  wrapper.innerHTML = [
    `<div class="alert theme-${type}" role="alert">`,
    `   <p>${message}</p>`,
    '   <button type="button" class="btn-close ms-auto" data-bs-dismiss="alert" aria-label="Close"></button>',
    '</div>'
  ].join('')

  liveAlertDiv.append(wrapper)
}

function handleBadRedirectStatus(status){

    switch (status) {
        case 400:
            //Bad request
            appendAlert("Bad Redirect. Please check requirements.", "warning");
            break;
        case 405:
            //Unauth
            appendAlert("Something went wrong.", "secondary");
            break;
            
        default:
            appendAlert("Server error: " + res.status, res.body, "warning");
            break;
    }
}