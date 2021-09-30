<template>
  <button :class="btnColor">
    <slot></slot>
  </button>
</template>

<script>
export default {
  props: ['btnColor'],
};
</script>

<style lang="scss" scoped>
  .btn {
    border: 0;
    cursor: pointer;
    box-sizing: border-box;
    border-radius: 3px;
    color: #fff;
    display: block;
    justify-content: space-around;
    position: relative;
  }
  .btn-pagination {
    background-color: rgb(175, 170, 170);
    border:1px solid rgb(126, 125, 125);
    padding:8px;
    margin: 80px;
  }
  .btn-small {
    padding: .5em 2em;
  }
  .btn-medium {
    padding: .7em 4.5em;
  }
  .btn-large {
    padding: .7em 5em;
  }
  .btn-info {
    background-color: transparent;
  }

  .btn-common {
    background-color:rgba(17, 64, 153, 0.897)

  }
  .btn-sucess {
    background-color: #27AE60;
  }
  .btn-danger {
    background-color: #e74c3c;
  }
  .btn img {
    max-width: 20px;
    margin-left: .5em;
  }

</style>