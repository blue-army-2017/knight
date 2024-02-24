const confirmElements = document.querySelectorAll(
  "a[data-confirm]:not([data-confirm=''])",
);

for (const element of confirmElements) {
  const url = element.getAttribute("href");
  const prompt = element.getAttribute("data-confirm");

  element.addEventListener("click", (e) => {
    e.preventDefault();

    if (window.confirm(prompt)) {
      window.location.assign(url);
    }
  });
}
