<template>
  <header>
    <nav id="navbar">
      <div class="container">
        <h1 class="logo"><a href="#">Forum</a></h1>
        <ul>
          <li><input type="text"></li>
          <li><a href="#">Search</a></li>
        </ul>
      </div>
    </nav>
  </header>
</template>

<script>
  export default {
    name: "Navbar"
  }
</script>

<style>
  .container {
    margin: auto;
    max-width: 1100px;
    overflow: auto;
    padding: 0 20px;
  }

  #navbar {
    background: #333;
    color: #fff;
    overflow: auto;
  }

  #navbar a {
    color: #fff;
  }

  #navbar .logo {
    float: left;
    padding-top: 20px;
  }

  #navbar ul {
    list-style: none;
    float: right;
  }

  #navbar ul li {
    float: left;
  }

  #navbar ul li a {
    display: block;
    padding: 20px;
    text-align: center;
  }

  #navbar ul li input {
    max-height: 40px;
    display: block;
    margin-top: 15px;
    padding: 20px;
    text-align: center;
  }

  #navbar ul li a:hover {
    background: #444;
    color: #f7c08a;
  }
</style>