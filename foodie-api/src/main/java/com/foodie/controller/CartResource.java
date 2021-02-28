package com.foodie.controller;

import com.foodie.dto.CartDto;
import java.util.Map;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/carts")
@Slf4j
public class CartResource {

  @PostMapping
  public ResponseEntity<Map<String, Object>> saveCart(@RequestBody CartDto cartDto) {
    log.info("saving cart {} ", cartDto.getCartItemDtoList());
    return ResponseEntity.ok(Map.of("status", "OK"));
  }
}
