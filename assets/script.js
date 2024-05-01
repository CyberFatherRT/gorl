const regex = /https?:\/\/[a-zA-Z0-9%]*:?[a-zA-Z0-9%]*\/?[(a-z).\/?]*\/?[^\s]+\.[^\s]{2,}\/?.*/;

async function create_random_link(link) {
    return await fetch("/api/v1/create_random_link", {
        method: "POST",
        body: JSON.stringify({
            long_link: link,
        }),
    }).then((res) => res.json());
}

async function shorten() {
    let enter_input = document.getElementById("enter-input");
    let link_input = document.getElementById("link-input");

    let link = enter_input.value;
    if (!regex.test(link)) {
        alert("Entered url is not valid!");
        return;
    }

    let result = await create_random_link(link);
    link_input.value = result.short_link;
}
