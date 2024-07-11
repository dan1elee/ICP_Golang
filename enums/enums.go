package enums

const STUDENT uint = 1
const TEACHER uint = (1 << 1) | 1
const ADMIN uint = (1 << 2) | (1 << 1) | 1

const LEVELNORMAL uint = 1
const LEVELSECRET uint = 1 << 1
const LEVELTOP uint = 1 << 2
