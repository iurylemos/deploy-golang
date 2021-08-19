const listMessages = document.querySelector(".chat-messages-list");
const chatInput = document.querySelector(".chat-input");
const messages = [];
let input = "";
let context = null;

document.querySelector(".chat-send").addEventListener("click", (evt) => {
  evt.preventDefault();
  sendMessage();
});

chatInput.addEventListener("input", (evt) => {
  evt.preventDefault();
  input = evt.target.outerText;
});

function bubbleUser(text) {
  const li = document.createElement("li");
  li.className = "chat-message chat-message-self";
  const bubble = document.createElement("div");
  bubble.className = "chat-message-bubble";
  bubble.innerText = text;
  li.appendChild(bubble);
  return li;
}

function bubbleChatbot(text) {
  const li = document.createElement("li");
  li.className = "chat-message chat-message-friend";
  const bubble = document.createElement("div");
  bubble.className = "chat-message-bubble";
  bubble.innerText = text;
  bubble.style = "transform: translate3d(0px, 0px, 0px); opacity: 1;";
  li.appendChild(bubble);
  return li;
}

function sendInput() {
  if (context) {
    return {
      input,
      context,
    };
  }

  return {
    input,
  };
}

async function sendMessage() {
  try {
    const createBuble = bubbleUser(input);

    const resp = await fetch("http://localhost:5000/api/messages", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify(sendInput()),
    });
    const json = await resp.json();

    // console.log("JSON", json);
    context = json.context;
    listMessages.appendChild(createBuble);

    if (json.output.text) {
      listMessages.appendChild(bubbleChatbot(json.output.text[0]));
    }

    chatInput.innerHTML = "";
  } catch (error) {
    console.log("error", JSON.parse(JSON.stringify(error)));
  }
}

(() => {
  listMessages.appendChild(bubbleChatbot("Bom dia"));
  //   listMessages.appendChild(bubbleUser("Bom dia"));
})();
