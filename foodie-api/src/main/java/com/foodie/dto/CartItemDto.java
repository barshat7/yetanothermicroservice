package com.foodie.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;

@NoArgsConstructor
@AllArgsConstructor
@Getter
public class CartItemDto {

  @JsonProperty("id")
  private Long itemID;

  @JsonProperty("name")
  private String itemName;

  @JsonProperty("price")
  private Double itemPrice;

  private int quantity;
}
