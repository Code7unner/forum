<template>
  <div class="topic-form-container">
    <Topic v-for="topic in topics" />
  </div>
</template>

<script>
import Topic from "./Topic";

export default {
  name: "TopicForm",
  static: {
    visibleItemsPerPageCount: 2
  },
  components: {
    Topic
  },
  data() {
    return {
      topics: 10,
      currentPage: 1,
      pageCount: 0
    };
  }
};
</script>

<style>
.topic-form-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
}
</style>
