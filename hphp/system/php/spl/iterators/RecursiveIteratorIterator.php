<?php

// This doc comment block generated by idl/sysdoc.php
/**
 * ( excerpt from
 * http://php.net/manual/en/class.recursiveiteratoriterator.php )
 *
 * Can be used to iterate through recursive iterators.
 *
 */
class RecursiveIteratorIterator implements OuterIterator {

  const LEAVES_ONLY = 0;
  const SELF_FIRST = 1;
  const CHILD_FIRST = 2;
  const CATCH_GET_CHILD = 16;

  const STATE_NEXT = 10;
  const STATE_TEST = 11;
  const STATE_SELF = 12;
  const STATE_CHILD = 13;
  const STATE_START = 14;

  const NEXT_COMPLETE = 10;
  const NEXT_REPEAT = 11;

  private $iterators = array();
  private $originalIterator;
  private $mode;
  private $flags;
  private $maxDepth = -1;

  // Flag for checking whether to call
  // begin/endIteration() or not
  private $inIteration = false;

  // This doc comment block generated by idl/sysdoc.php
  /**
   * ( excerpt from
   * http://php.net/manual/en/recursiveiteratoriterator.construct.php )
   *
   * Creates a RecursiveIteratorIterator from a RecursiveIterator.
   *
   * @iterator   mixed   The iterator being constructed from. Either a
   *                     RecursiveIterator or IteratorAggregate.
   * @mode       mixed   Optional mode. Possible values are
   *                     RecursiveIteratorIterator::LEAVES_ONLY - The
   *                     default. Lists only leaves in iteration.
   *                     RecursiveIteratorIterator::SELF_FIRST - Lists leaves
   *                     and parents in iteration with parents coming first.
   *                     RecursiveIteratorIterator::CHILD_FIRST - Lists
   *                     leaves and parents in iteration with leaves coming
   *                     first.
   * @flags      mixed   Optional flag. Possible values are
   *                     RecursiveIteratorIterator::CATCH_GET_CHILD which
   *                     will then ignore exceptions thrown in calls to
   *                     RecursiveIteratorIterator::getChildren().
   *
   * @return     mixed   No value is returned.
   */
  public function __construct(\Traversable $iterator,
                              $mode = RecursiveIteratorIterator::LEAVES_ONLY,
                              $flags = 0) {
    if ($iterator && ($iterator is IteratorAggregate)) {
      $iterator = $iterator->getIterator();
    }
    if (!$iterator || !($iterator is RecursiveIterator)) {
      throw new InvalidArgumentException(
        "An instance of RecursiveIterator or IteratorAggregate creating " .
        "it is required"
      );
    }
    $this->iterators[] = array($iterator, self::STATE_START);
    $this->originalIterator = $iterator;
    $this->mode = (int) $mode;
    $this->flags = $flags;
  }

  // This doc comment block generated by idl/sysdoc.php
  /**
   * ( excerpt from
   * http://php.net/manual/en/recursiveiteratoriterator.getinneriterator.php
   * )
   *
   * Gets the current active sub iterator. Warning: This function is
   * currently not documented; only its argument list is available.
   *
   * @return     mixed   The current active sub iterator.
   */
  public function getInnerIterator() {
    $it = $this->iterators[count($this->iterators)-1][0];
    if (!$it is RecursiveIterator) {
      throw new Exception(
        "inner iterator must implement RecursiveIterator"
      );
    }
    return $it;
  }

  // This doc comment block generated by idl/sysdoc.php
  /**
   * ( excerpt from
   * http://php.net/manual/en/recursiveiteratoriterator.current.php )
   *
   *
   * @return     mixed   The current elements value.
   */
  public function current() {
    return $this->getInnerIterator()->current();
  }

  // This doc comment block generated by idl/sysdoc.php
  /**
   * ( excerpt from
   * http://php.net/manual/en/recursiveiteratoriterator.key.php )
   *
   *
   * @return     mixed   The current key.
   */
  public function key() {
    return $this->getInnerIterator()->key();
  }

  // This doc comment block generated by idl/sysdoc.php
  /**
   * ( excerpt from
   * http://php.net/manual/en/recursiveiteratoriterator.next.php )
   *
   *
   * @return     mixed   No value is returned.
   */
  public function next() {
    $this->nextImpl();
  }

  private function nextImpl() {
    while ($this->nextInnerImpl() === self::NEXT_REPEAT) {
      /* loop */
    }
  }

  /* We maintain a stack of sub-iterators - each of which has a state.
   *
   * This walks the overall tree for one step, and updates the corresponding
   * iterator states as needed (current iterator, parent iterator, child
   * iterator).
   *
   * Returns self::NEXT_COMPLETE if that one step got us to the next position
   *   (this changes depending on LEAVES_ONLY, SELF_FIRST, and CHILD_FIRST), or
   *   SELF::NEXT_REPEAT if more steps are needed.
   */
  private function nextInnerImpl() {
    $it = $this->getInnerIterator();
    switch ($this->getInnerIteratorState()) {
      case self::STATE_NEXT:
        $it->next();
        // fallthrough
      case self::STATE_START:
        if (!$it->valid()) {
          if ($this->getDepth() > 0) {
            $this->endChildren();
            array_pop(&$this->iterators);
            return self::NEXT_REPEAT;
          }
          return self::NEXT_COMPLETE;
        }
        $this->setInnerIteratorState(self::STATE_TEST);
        return self::NEXT_REPEAT;
      case self::STATE_TEST:
        if ($this->callHasChildren()) {
          if ($this->maxDepth == -1 || $this->maxDepth > $this->getDepth()) {
            switch ($this->mode) {
              case self::LEAVES_ONLY:
              case self::CHILD_FIRST:
                // We never look at SELF in LEAVES_ONLY
                $this->setInnerIteratorState(self::STATE_CHILD);
                return self::NEXT_REPEAT;
              case self::SELF_FIRST:
                $this->setInnerIteratorState(self::STATE_SELF);
                return self::NEXT_REPEAT;
            }
          } else if ($this->mode == self::LEAVES_ONLY) {
            // We're already at the recursion limit, and the current node isn't
            // a leaf
            $this->setInnerIteratorState(self::STATE_NEXT);
            return self::NEXT_REPEAT;
          }
        }
        $this->nextElement();
        $this->setInnerIteratorState(self::STATE_NEXT);
        return self::NEXT_COMPLETE;
      case self::STATE_SELF:
        $this->nextElement();
        if ($this->mode == self::SELF_FIRST) {
          $this->setInnerIteratorState(self::STATE_CHILD);
        } else {
          $this->setInnerIteratorState(self::STATE_NEXT);
        }
        return self::NEXT_COMPLETE;
      case self::STATE_CHILD:
        try {
          $children = $this->callGetChildren();
        } catch (Exception $e) {
          if (!($this->flags == self::CATCH_GET_CHILD)) throw $e;
          $this->setInnerIteratorState(self::STATE_NEXT);
          return self::NEXT_REPEAT;
        }
        if (!$children is RecursiveIterator) {
          throw new UnexpectedValueException(
            'Objects returned by RecursiveIterator::getChildren() must '.
            'implement RecursiveIterator'
          );
        }
        if ($this->mode == self::CHILD_FIRST) {
          $this->setInnerIteratorState(self::STATE_SELF);
        } else {
          $this->setInnerIteratorState(self::STATE_NEXT);
        }
        $children->rewind();
        $this->iterators[] = array($children, self::STATE_START);
        $this->beginChildren();
        return self::NEXT_REPEAT;
    }
    throw new Exception("unreachable");
  }

  // This doc comment block generated by idl/sysdoc.php
  /**
   * ( excerpt from
   * http://php.net/manual/en/recursiveiteratoriterator.rewind.php )
   *
   *
   * @return     mixed   No value is returned.
   */
  public function rewind() {
    while ($this->iterators) {
      array_pop(&$this->iterators);
    }

    $it = $this->originalIterator;
    $this->iterators = array(array($it, self::STATE_START));
    $it->rewind();
    if (!$this->inIteration) {
      $this->beginIteration();
    }

    $this->nextImpl();
    $this->inIteration = true;
  }

  // This doc comment block generated by idl/sysdoc.php
  /**
   * ( excerpt from
   * http://php.net/manual/en/recursiveiteratoriterator.valid.php )
   *
   *
   * @return     mixed   TRUE if the current position is valid, otherwise
   *                     FALSE
   */
  public function valid() {
    $depth = $this->getDepth();
    while ($depth >= 0) {
      if ($this->getSubIterator($depth)->valid()) {
        return true;
      }
      $depth--;
    }
    if ($this->inIteration) {
      $this->endIteration();
    }
    $this->inIteration = false;
    return false;
  }

  /**
   * Called after calling getChildren(), and its associated rewind().
   */
  public function beginChildren()
  {
  }

  /**
   * Called when iteration begins (after the first rewind() call).
   */
  public function beginIteration()
  {
  }

  /**
   * Get children of the current element.
   *
   * @return     RecursiveIterator
   */
  public function callGetChildren()
  {
    return $this->getInnerIterator()->getChildren();
  }

  /**
   * Called for each element to test whether it has children.
   *
   * @return     bool
   */
  public function callHasChildren()
  {
    return $this->getInnerIterator()->hasChildren();
  }

  /**
   * Called when end recursing one level.
   */
  public function endChildren()
  {
  }

  /**
   * Called when the iteration ends (when valid() first returns FALSE).
   */
  public function endIteration()
  {
  }

  /**
   * Get the current depth of the recursive iteration.
   *
   * @return     int     The current depth of the recursive iteration.
   */
  public function getDepth()
  {
    return count($this->iterators)-1;
  }

  /**
   * Gets the maximum allowable depth.
   *
   * @return     int     The maximum accepted depth, or FALSE if any depth is
   *                     allowed.
   */
  public function getMaxDepth()
  {
    return ($this->maxDepth == -1) ? false : $this->maxDepth;
  }

  /**
   * Gets the current active sub iterator.
   *
   * @param      int     $level
   *
   * @return     RecursiveIterator   The current active sub iterator.
   */
  public function getSubIterator($level = null)
  {
    $currentLevel = count($this->iterators)-1;
    if (is_null($level)) {
      $level = $currentLevel;
    }
    if ($level < 0 || $level > $currentLevel) {
      return null;
    }
    return $this->iterators[$level][0];
  }

  /**
   * Called when the next element is available.
   */
  public function nextElement()
  {
  }

  /**
   * Set the maximum allowed depth.
   *
   * @param      int     $max_depth   The maximum allowed depth. -1 is used for
   *                                  any depth.
   *
   * @throws     Exception            Emits an Exception if max_depth is less
   *                                  than -1.
   */
  public function setMaxDepth($max_depth = -1)
  {
    if ($max_depth < -1) {
      throw new OutOfRangeException("Parameter max_depth must be >= -1");
    }

    $this->maxDepth = $max_depth;
  }

  private function isEmpty() {
    return count($this->iterators) == 0;
  }

  private function getInnerIteratorState() {
    return $this->iterators[count($this->iterators)-1][1];
  }

  private function setInnerIteratorState($state) {
    $this->iterators[count($this->iterators)-1][1] = $state;
  }

  /**
   * Undocumented behavior but Zend does it and frameworks rely on it, so..
   */
  public function __call($func, $params) {
    return call_user_func_array(
      array($this->getInnerIterator(), $func),
      $params
    );
  }
}
