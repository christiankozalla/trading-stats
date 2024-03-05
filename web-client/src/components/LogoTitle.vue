<template>
  <div class="logo-container">
    <img class="logo-image" src="@/assets/inloopo-logo.png" alt="Inloopo logo" />
    <h1 class="logo-text">Trade&nbsp;Stats</h1>
  </div>
</template>

<style scoped>
.logo-container {
  display: flex;
  height: 100%;

  align-items: center;
  gap: var(--inline-spacing);
}

.logo-image {
  aspect-ratio: 1;
  width: 32px;
}

.logo-text {
  display: inline-block;
  font-size: large;
}

@media screen and (max-width: 320px) {
  .logo-text {
    display: none;
  }
}
</style>
