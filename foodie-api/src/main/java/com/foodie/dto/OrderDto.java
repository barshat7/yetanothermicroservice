package com.foodie.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.ToString;

@NoArgsConstructor
@AllArgsConstructor
@Getter
@Builder
@ToString
public class OrderDto {

  private Long id;

  @JsonProperty("user_id")
  private String userID;

  @JsonProperty("order_date")
  private String orderDate;

  @JsonProperty("total_price")
  private double totalPrice;
}
