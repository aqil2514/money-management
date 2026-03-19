<template>
  <div>
    <h1>Profil Pengguna</h1>
    <p>ID Pengguna yang sedang diakses: {{ $route.params.id }}</p>

    <NuxtLink to="/">Kembali ke Beranda</NuxtLink>
  </div>
</template>