const listMessages = document.querySelector(".chat-messages-list");
const chatInput = document.querySelector(".chat-input");
const loading = document.getElementById("loading");
const buttonSend = document.querySelector(".chat-send");
const chatMessages = document.querySelector(".chat-messages");
const messages = [];
let input = "";
let context = null;
loading.style.display = "none";

buttonSend.addEventListener("click", (evt) => {
  evt.preventDefault();
  sendMessage();
});

chatInput.addEventListener("input", (evt) => {
  evt.preventDefault();
  input = evt.target.outerText;
});

chatInput.addEventListener("keydown", (evt) => {
  // evt.preventDefault();
  if (evt.key === "Enter") {
    evt.preventDefault();
    input = evt.target.outerText;
    sendMessage();
    // Do more work
  }
});

function scrollToBottom() {
  chatMessages.scrollTop = chatMessages.scrollHeight;
}

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

function setLoad(val) {
  if (val) {
    loading.style.display = "flex";
    buttonSend.setAttribute("disabled", true);
    chatInput.setAttribute("disabled", true);
  } else {
    loading.style.display = "none";
    buttonSend.removeAttribute("disabled");
    chatInput.removeAttribute("disabled");
  }
}

async function sendMessage() {
  try {
    chatInput.innerHTML = "";
    const createBuble = bubbleUser(input);
    setLoad(true);

    const resp = await fetch(`${window.location.origin}/api/messages`, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify(sendInput()),
    });
    const json = await resp.json();

    console.log("JSON", json);
    setLoad(false);
    context = json.context;
    listMessages.appendChild(createBuble);

    if (json.output.text) {
      listMessages.appendChild(bubbleChatbot(json.output.text[0]));
    }
    scrollToBottom();
  } catch (error) {
    console.log("error", error);
  }
}

(() => {
  const today = new Date();
  const curHr = today.getHours();

  if (curHr < 12) {
    listMessages.appendChild(bubbleChatbot("Bom dia"));
  } else if (curHr < 18) {
    listMessages.appendChild(bubbleChatbot("Boa tarde"));
  } else {
    listMessages.appendChild(bubbleChatbot("Boa noite"));
  }
})();
