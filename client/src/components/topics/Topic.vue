<template>
  <div class="topic-container">
    <div class="name">ШОККК САБАКУ ИБАЛИ В ЖЕПУУ С СМС И РЕГОЙ ПОД СОЛЯМИ</div>
    <div class="total-users">СКОЛЬКО ЮЗЕРОВ: ДОХУЯ</div>
    <div class="total-messages">СКОЛЬКО СООБЩЕНИ: НИХУЯ</div>
    <div class="last-message">ПОСЛЕДНЕЕ СООБЩЕНИЕ: ХУЙ ЗНАЕТ</div>
    <div class="author-name">МАКС ЛОХ</div>
  </div>
</template>

<script>
export default {
  name: "Topic"
};
</script>

<style lang="scss">
.topic-container {
  display: grid;
  width: 80%;
  grid-template-columns: repeat(12, 1fr);
  grid-template-areas:
    "n n n n n n tu tu tu lm lm lm"
    "n n n n n n tm tm tm an an an";
  grid-auto-rows: 40px;

  .name {
    grid-area: n;
    display: flex;
    justify-content: flex-start;
    align-item: center;
    font-size: 16px;
  }

  .total-users {
    grid-area: tu;
    font-size: 12px;
    display: flex;
    justify-content: center;
    align-item: center;
  }

  .total-messages {
    grid-area: tm;
    font-size: 12px;
    display: flex;
    justify-content: center;
    align-item: center;
  }

  .last-message {
    font-size: 12px;
    grid-area: lm;
    display: flex;
    justify-content: center;
    align-item: center;
  }

  .author-name {
    font-size: 12px;
    display: flex;
    justify-content: center;
    align-item: center;
    grid-area: an;
  }
}
</style>
