package com.foodie.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;

@NoArgsConstructor
@AllArgsConstructor
@Getter
public class CartItemDto {

  private Long itemID;

  private String itemName;

  private Double itemPrice;

  private int quantity;
}
