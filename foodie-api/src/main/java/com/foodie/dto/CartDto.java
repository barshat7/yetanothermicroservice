package com.foodie.dto;

import java.util.ArrayList;
import java.util.List;
import lombok.Data;

@Data
public class CartDto {

  private String userID;

  private List<CartItemDto> cartItemDtoList = new ArrayList<>();
}
