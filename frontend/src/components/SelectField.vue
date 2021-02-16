<template>
  <div class="select-wrapper">
    <label>{{ data.data.label }}</label>
    <v-select :options="data.data.options" @input="setSelected"/>
  </div>
</template>

<script>
export default {
  name: "SelectField",
  props: ['data'],
  methods: {
    setSelected(value) {
      console.log('setSelected', value);
      this.$emit('getValue', {id: this.$props.data.id, label: value.label});
    }
  }
}
</script>

<style>
.select-wrapper {
  margin-top: 15px;
  display: flex;
  flex-direction: column;
}
.select-wrapper label {
  font-weight: 500;
  text-align: left;
}
.vs__search {
  width: 100%;
  height: 32px;
  background-color: #ffffff;
  border: 1px solid #6F8BE9;
  border-radius: 4px;
  transition: 180ms box-shadow ease-in-out;
  font-size: 16px;
  font-family: inherit;
  padding: 0.25em 0.5em;
}
.vs__actions {
  position: absolute;
  top: 5px;
  left: 94%;
}
.vs__dropdown-toggle {
  position: relative;
}
.vs__selected {
  position: absolute;
  height: 32px;
  display: flex;
  justify-content: center;
  align-items: center;
  padding-left: 9px;
}
.vs__clear {
  position: absolute;
  right: 25px;
}
</style>
