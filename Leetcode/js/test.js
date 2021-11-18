const test = [
  {
    id: "01",
    shopName: "nutShop",
    categories: "food",
    province: "aran",
  },
  {
    id: "02",
    shopName: "patShop",
    categories: "food",
    province: "aran",
  },
  {
    id: "03",
    shopName: "maiShop",
    categories: "fashion",
    province: "aran",
  },
  {
    id: "04",
    shopName: "armShop",
    categories: "fashion",
    province: "aran",
  },
  {
    id: "05",
    shopName: "tosShop",
    categories: "name",
    province: "aran",
  },
];

const a = test.filter((item) => item.province === "aran");
