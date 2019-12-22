<template>
  <div>
    <nav-bar></nav-bar>
    <topic-form></topic-form>
  </div>
</template>

<script>
  import NavBar from "./navbar/Navbar"
  import TopicForm from "./topics/TopicForm";

  export default {
    name: "Home",
      components: {
        NavBar,
        TopicForm
      },
      data() {
        return {}
      }
  }
</script>

<style>

</style>