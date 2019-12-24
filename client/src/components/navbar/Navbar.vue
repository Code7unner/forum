<template>
  <header>
    <nav id="navbar">
      <div class="navbar-container">
        <h1 class="logo"><a href="#">Forum</a></h1>
        <ul>
          <li><input type="text" /></li>
          <li><a href="#">Search</a></li>
        </ul>
      </div>
    </nav>
  </header>
</template>

<script>
export default {
  name: "Navbar"
};
</script>

<style lang="scss">
.navbar-container {
  margin: auto;
  max-width: 1100px;
  overflow: auto;
  padding: 0 20px;
}

#navbar {
  background: #333;
  color: #fff;
  overflow: auto;

  a {
    color: #fff;
  }

  .logo {
    float: left;
    padding-top: 20px;
  }

  ul {
    list-style: none;
    float: right;

    li {
      float: left;

      a {
        display: block;
        padding: 20px;
        text-align: center;
        transition: 0.1s;

        &:hover {
          background: #444;
          color: #f7c08a;
        }
      }

      input {
        max-height: 40px;
        display: block;
        margin-top: 15px;
        padding: 20px;
        text-align: center;
      }
    }
  }
}
</style>
